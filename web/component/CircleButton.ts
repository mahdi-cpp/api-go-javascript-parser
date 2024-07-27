import type {ViewProps, ViewValues} from "./view";
import {ThemValues, TouchProps} from "./view";

export interface CircleButtonProps extends ViewProps, TouchProps {
    icon?: string,
    checked?: boolean,
    duration?: number,
    line?: boolean,
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

export type CircleButtonValues = ViewValues & ThemValues & {
    icon: string,
    checked: number;
    duration: number;
    line: number;
};

export default function CircleButton(props: CircleButtonProps) {

}


