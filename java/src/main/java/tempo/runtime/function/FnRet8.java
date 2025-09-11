package tempo.runtime.function;

@FunctionalInterface
public interface FnRet8<Return, One, Two, Three, Four, Five, Six, Seven, Eight> {
    public Return call(One one, Two two, Three three, Four four, Five five, Six six, Seven seven, Eight eight)
            throws Exception;
}
