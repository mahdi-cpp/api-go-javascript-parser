import type {ViewProps, ViewValues} from "../view";
import {ThemValues, TouchProps} from "../view";

export interface ChartViewProps extends ViewProps, TouchProps {

    backgroundColor?: string,
    activeColor?: string,

    headerHeight?: number,
    chartHeight?: number,
    footerHeight?: number,

    avatar: string
    title: string
    caption?: string

    dataArray: number[],
    rowArray: string[],
    columnArray: string[],
}

export type ChartViewValues = ViewValues & ThemValues & {
    backgroundColor?: string,
    activeColor?: string,
    circleColor?: string,

    headerHeight?: number,
    chartHeight?: number,
    footerHeight?: number,

    title: string
    caption?: string

    rowCount: number,
    columnCount: number,
};

export default function ChartView(props: ChartViewProps) {

}


