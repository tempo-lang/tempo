package tempo.runtime.function;

@FunctionalInterface
public interface FnRet9<Return, One, Two, Three, Four, Five, Six, Seven, Eight, Nine> {
    public Return call(One one, Two two, Three three, Four four, Five five, Six six, Seven seven, Eight eight,
            Nine nine) throws Exception;
}
