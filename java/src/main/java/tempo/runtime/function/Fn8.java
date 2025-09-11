package tempo.runtime.function;

@FunctionalInterface
public interface Fn8<One, Two, Three, Four, Five, Six, Seven, Eight> {
    public void call(One one, Two two, Three three, Four four, Five five, Six six, Seven seven, Eight eight)
            throws Exception;
}
