import type {TextProps, TextValues, ViewProps, ViewValues} from "../view";
import {ThemValues, TouchProps} from "../view";

export interface TextBoxProps extends ViewProps, TouchProps, TextProps {
    boxColor?: string
}

export type TextBoxValues = ViewValues & ThemValues & TextValues & {
    boxColor?: string
};

export default function TextBox(props: TextBoxProps) {

}


