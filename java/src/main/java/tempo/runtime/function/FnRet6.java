package tempo.runtime.function;

@FunctionalInterface
public interface FnRet6<Return, One, Two, Three, Four, Five, Six> {
    public Return call(One one, Two two, Three three, Four four, Five five, Six six) throws Exception;
}
