import type {ViewProps, ViewValues} from "../view";
import {TextProps, TextValues, ThemValues, TouchProps} from "../view";

export interface ScrollViewProps extends ViewProps, TouchProps, TextProps {
    icon?: string,

    backgroundColor?: string,
    activeColor?: string,
    circleColor?: string,

    // onChange(): void;
    //
    // play(): void;
    //
    // pause(): void;
    //
    // next(url: string)
}

export type ScrollViewValues = ViewValues & ThemValues & TextValues & {
    icon: string,
    checked: number;
    duration: number;
    line: number;
};

export default function ScrollView(props: ScrollViewProps) {

}


