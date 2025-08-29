package tempo.runtime.function;

@FunctionalInterface
public interface FnRet4<Return, One, Two, Three, Four> {
    public Return call(One one, Two two, Three three, Four four) throws Exception;
}
