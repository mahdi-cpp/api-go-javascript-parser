import type {ViewProps, ViewValues} from "../view";
import {ThemValues, TouchProps} from "../view";

export interface SliderViewProps extends ViewProps, TouchProps {
    icon?: string,
    checked?: boolean,
    duration?: number,
    line?: boolean,
    backgroundColor?: string,
    activeColor?: string,
    circleColor?: string,
    circleSize?: string,
}

export type SliderViewValues = ViewValues & ThemValues & {
    icon: string,
    checked: number;
    duration: number;
    line: number;
    circleSize?: string,
};

export default function SliderView(props: SliderViewProps) {

}


