package tempo.runtime.function;

@FunctionalInterface
public interface FnRet5<Return, One, Two, Three, Four, Five> {
    public Return call(One one, Two two, Three three, Four four, Five five) throws Exception;
}
