import type {TextProps, TextValues, ViewProps, ViewValues} from "../view";
import {ThemValues, TouchProps} from "../view";

export interface TextMessageProps extends ViewProps, TouchProps, TextProps {
    boxColor?: string
}

export type TextMessageValues = ViewValues & ThemValues & TextValues & {
    boxColor?: string
};

export default function TextMessage(props: TextMessageProps) {

}


