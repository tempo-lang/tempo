package tempo.runtime.function;

@FunctionalInterface
public interface FnRet1<Return, One> {
    public Return call(One one) throws Exception;
}
