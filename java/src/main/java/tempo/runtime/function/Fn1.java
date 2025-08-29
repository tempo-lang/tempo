package tempo.runtime.function;

@FunctionalInterface
public interface Fn1<One> {
    public void call(One one) throws Exception;
}
