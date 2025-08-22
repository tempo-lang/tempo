package tempo.runtime.function;

@FunctionalInterface
public interface FnRet2<Return, One, Two> {
    public Return call(One one, Two two);
}
