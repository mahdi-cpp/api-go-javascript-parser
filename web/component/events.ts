class EventHandler<T> {
}

class SyntheticEvent<T> {
}

type TouchEventHandler<T = Element> = EventHandler<SyntheticEvent<T>>;