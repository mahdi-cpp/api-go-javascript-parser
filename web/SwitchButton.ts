import type {ViewProps, ViewValues} from "./view";
import {ThemValues, TouchProps} from "./view";

export interface SwitchButtonProps extends ViewProps, TouchProps {
    icon?: string,
    checked?: boolean,
    duration?: number,
    line?: boolean,
    backgroundColor?: string,
    activeColor?: string,
    circleColor?: string,

    onChange(): void;

    play(): void;

    pause(): void;

    next(url: string)
}

export type SwitchButtonValues = ViewValues & ThemValues & {
    icon: string,
    checked: number;
    duration: number;
    line: number;
};

export default function SwitchButton(props: SwitchButtonProps) {

}


