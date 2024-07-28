import type {ViewProps, ViewValues} from "../view";
import {TextProps, TextValues, ThemValues, TouchProps} from "../view";

export interface ButtonProps extends ViewProps, TouchProps, TextProps {
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

export type ButtonValues = ViewValues & ThemValues & TextValues & {
    icon: string,
    checked: number;
    duration: number;
    line: number;
};

export default function Button(props: ButtonProps) {

}


