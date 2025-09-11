package tempo.runtime.function;

@FunctionalInterface
public interface Fn6<One, Two, Three, Four, Five, Six> {
    public void call(One one, Two two, Three three, Four four, Five five, Six six) throws Exception;
}
