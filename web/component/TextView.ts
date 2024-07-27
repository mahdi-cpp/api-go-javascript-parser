import type {TextProps, TextValues, ViewProps, ViewValues} from "./view";
import {ThemValues, TouchProps} from "./view";

export interface TextViewProps extends ViewProps, TouchProps, TextProps {

}

export type TextViewValues = ViewValues & ThemValues & TextValues & {};

export default function TextView(props: TextViewProps) {

}


