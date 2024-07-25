import {TouchEventHandler} from "./events";

export interface ViewProps {
    id?: string;
    dx?: boolean,
    dy?: number,
    width?: number,
    height?: number,
    round?: number,
}

export interface TouchProps {
    onClick?: TouchEventHandler,
    onPinch?: TouchEventHandler,
}

export type ViewValues = {
    id: number;
    dx: number;
    dy: number;
    width: number;
    height: number;
    round: number;
};

export type ThemValues = {
    backgroundColor: number;
    activeColor: number;
    circleColor: number;
};












