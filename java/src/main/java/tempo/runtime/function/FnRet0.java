package tempo.runtime.function;

@FunctionalInterface
public interface FnRet0<Return> {
    public Return call() throws Exception;
}
