package tempo.runtime.function;

@FunctionalInterface
public interface FnRet7<Return, One, Two, Three, Four, Five, Six, Seven> {
    public Return call(One one, Two two, Three three, Four four, Five five, Six six, Seven seven) throws Exception;
}
