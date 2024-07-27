import {TouchEventHandler} from "./events";

export interface ViewProps {
    id?: string;
    dx?: boolean,
    dy?: number,
    width?: number,
    height?: number,
    round?: number,
    padding?: number,
    margin?: number,

    icon?: string,

}

export interface TouchProps {
    onClick?: TouchEventHandler,
    onPinch?: TouchEventHandler,
}

export interface TextProps {
    backgroundColor?: string,
    textColor?: string,
    text: string,
    textSize?: string,
}

export type ViewValues = {
    id: number;
    dx: number;
    dy: number;
    width: number;
    height: number;
    round: number;
    padding?: number,
};

export type ThemValues = {
    backgroundColor: number;
    activeColor: number;
    circleColor: number;
};

export type TextValues = {
    backgroundColor?: string,
    textColor?: string,
    text: string,
    textSize?: string,
};











