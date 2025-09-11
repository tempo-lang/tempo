package tempo.runtime.function;

@FunctionalInterface
public interface Fn7<One, Two, Three, Four, Five, Six, Seven> {
    public void call(One one, Two two, Three three, Four four, Five five, Six six, Seven seven) throws Exception;
}
