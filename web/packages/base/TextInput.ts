import type {TextProps, TextValues, ViewProps, ViewValues} from "../view";
import {ThemValues, TouchProps} from "../view";

export interface TextInputProps extends ViewProps, TouchProps, TextProps {
    boxColor?: string
}

export type TextInputValues = ViewValues & ThemValues & TextValues & {
    boxColor?: string
};

export default function TextInput(props: TextInputProps) {

}


