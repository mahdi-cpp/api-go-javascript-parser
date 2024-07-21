interface ImageProps {
    id?: string;
    source?: string;
    dx?: boolean,
    dy?: number,
    width?: number,
    height?: number,
    round?: number,

    onClick?: TouchEventHandler,
    onPinch?: TouchEventHandler,

}

export const ImageValues = {
    id: {},
    source: {},
    dx: {},
    dy: {},
    width: {},
    height: {},
    round: {},
};

export default function Image(props: ImageProps) {

}