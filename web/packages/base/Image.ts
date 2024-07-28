import {TouchProps, ViewProps, ViewValues} from "../view";

interface ImageProps extends ViewProps, TouchProps{
    source?: string;
}

export type ImageValues = ViewValues & {
    source: string;
};

export default function Image(props: ImageProps) {

}