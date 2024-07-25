class EventHandler<T> {
}

class SyntheticEvent<T> {
}

export type TouchEventHandler<T = Element> = EventHandler<SyntheticEvent<T>>;