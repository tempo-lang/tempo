package tempo.runtime.function;

@FunctionalInterface
public interface FnRet3<Return, One, Two, Three> {
    public Return call(One one, Two two, Three three);
}
