package tempo.runtime.function;

@FunctionalInterface
public interface Fn5<One, Two, Three, Four, Five> {
    public void call(One one, Two two, Three three, Four four, Five five) throws Exception;
}
