package tempo.simulator;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

import tempo.runtime.Env;
import tempo.transports.LocalQueue;
import tempo.transports.Recorder;

public class Simulator {

    protected HashMap<String, Process> processes;

    public Simulator() {
        this(Map.of());
    }

    public Simulator(Map<String, Process> processes) {
        this.processes = new HashMap<>(processes);
    }

    public Simulator addProcess(String role, Process process) {
        processes.put(role, process);
        return this;
    }

    public Map<String, Result> run() throws InterruptedException, ProcessExecutionException {
        LocalQueue queue = new LocalQueue();

        List<String> roles = new ArrayList<>(processes.keySet());

        List<Callable<Result>> tasks = roles.stream().map(role -> (Callable<Result>) () -> {
            Recorder transport = new Recorder(queue.role(role));
            Env env = new Env(transport);
            Object returnVal = processes.get(role).run(env);
            return new Result(returnVal, transport.sendValues(), transport.receivedValues());
        }).toList();

        ExecutorService executor = Executors.newCachedThreadPool();

        HashMap<String, Result> results = new HashMap<>();

        List<Future<Result>> futures = executor.invokeAll(tasks);

        for (int i = 0; i < roles.size(); i++) {
            var role = roles.get(i);
            try {
                var result = futures.get(i).get();
                results.put(role, result);
            } catch (ExecutionException e) {
                throw new ProcessExecutionException(role, e);
            }
        }

        return results;
    }
}
