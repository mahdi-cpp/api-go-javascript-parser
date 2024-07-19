function Menu() {
    return null;
}

function MenuIcon() {
    return null;
}

function MenuSwitch() {
    return null;
}

function MenuHeader() {
    return null;
}

function MenuTitle() {
    return null;
}

function View(props) {
    return null;
}

const MenuView = () => (
    <View>
        <Menu id={'more-menu'} >
            <MenuIcon
                title="Delete of library"
                icon={"icons/delete"}
            />
            <MenuSwitch
                title="Play music always"
                icon={"icons/music_play"}
            />
            <MenuHeader
                height={20}
            />
            <MenuTitle
                title="Settings"
            />
            <MenuIcon
                title="Download"
                icon={"icons/download"}
            />
        </Menu>
    </View>
);