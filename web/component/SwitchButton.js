interface SwitchButtonProps {
    id?: string;
    dx?: boolean,
    dy?: number,
    width?: number,
    height?: number,
    round?: number,

    checked?: boolean,
    duration?: number,
    line?: boolean,
    backgroundColor?: string,
    activeColor?: string,
    circleColor?: string,

    onClick?: TouchEventHandler,
    onPinch?: TouchEventHandler,
}

export const SwitchButtonValues  = {
    id: {},
    dx: {},
    dy: {},
    width: {},
    height: {},
    round: {},

    checked: {},
    duration: {},
    line: {},
    backgroundColor: {},
    activeColor: {},
    circleColor: {},
}

export default function SwitchButton(props: SwitchButtonProps) {

}