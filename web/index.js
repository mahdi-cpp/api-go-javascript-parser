import SwitchButton, {SwitchButtonValues} from "./component/SwitchButton";
import Image, {ImageValues} from "./component/Image";

function changePhoto() {

    let switch1 = SwitchButtonValues;
    let switch2 = SwitchButtonValues;
    let avatar = ImageValues;
    let avatar2 = ImageValues;

    switch1.id = 'switch1';
    switch1.dx = 645;
    switch1.checked = true;
    switch1.line = true;
    switch1.width = 400
    switch1.height = 200;
    switch1.round = 1;
    switch1.activeColor = '#ff9800'

    switch2.id = 'switch2';
    switch2.width = 350;
    switch2.height = 200
    switch2.dx = 695;
    switch2.dy = 550;
    switch2.round = 100;
    switch2.backgroundColor = '#009900';

    avatar.id = 'avatar2';
    avatar.source = 'music/face/peoples_86.jpg';
    avatar.dx = 150;
    avatar.dy = 200;
    avatar.width = 500;
    avatar.height = 500;
    avatar.round = 3;

    avatar2.id = 'avatar1';
    avatar2.source = 'music/face/peoples_73.jpg';
    avatar2.width = 500;
    avatar2.height = 500;
    avatar2.dx = 150;
    avatar2.dy = 1350;
    avatar2.round = 10;

}

function changePhoto2() {
    let image1 = ImageValues
    image1.id = 'image1'
    image1.source = 'call2/ali4.jpg'
}

const Index = () => (
    <View>
        <SwitchButton
            onClick={changePhoto2}
            onPinch={}
        />
        <SwitchButton
            id={'switchMain'}
            dx={720}
            dy={260}
            width={300}
            height={190}
            round={5}
            checked={false}
            duration={250}
            line={true}
            backgroundColor={'#eceff1'}
            activeColor={'#f44336'}
            circleColor={'#fff'}
            onClick={changePhoto}
        />
        <SwitchButton
            id={'switch2'}
            dx={820}
            dy={700}
            width={200}
            height={90}
            round={5}
            checked={false}
            duration={250}
            line={true}
            backgroundColor={'#eceff1'}
            activeColor={'#ff9800'}
            circleColor={'#fff'}
            onClick={changePhoto}
        />
        <SwitchButton
            id={'switch1'}
            icon={'icons/nav_check.png'}
            dx={730}
            dy={1200}
            width={250}
            height={120}
            round={30}
            checked={false}
            duration={200}
            line={false}
            backgroundColor={'#b0bec5'}
            activeColor={'#000'}
            circleColor={'#fff'}
            onClick={changePhoto2}
        />
        <Image
            id={"avatar1"}
            source={'call2/ali5.jpg'}
            dx={50}
            dy={210}
            width={170}
            height={170}
            round={55}
            onClick={changePhoto}
        />
        <Image
            id={"avatar2"}
            source={'call2/ali4.jpg'}
            dx={250}
            dy={210}
            width={170}
            height={170}
            round={55}
            onClick={changePhoto}
        />
        <TextView style={styles.title}
                  dx={240}
                  dy={280}
                  width={500}
                  height={200}
                  textSize={16}
                  textColor={'#000000'}
                  align={'NORMAL'}
                  text={'Farahmand Alipor'}
        />
        <SliderView
            dx={80}
            dy={1050}
            width={960}
            height={20}
            round={55}
            circleSize={10}
            line={true}
            backgroundColor={'#eee'}
            activeColor={'#55f'}
            circleColor={'#55f'}
        />
        <SliderView
            dx={80}
            dy={800}
            width={910}
            height={20}
            round={50}
            circleSize={150}
            line={true}
            backgroundColor={'#fbe9e7'}
            activeColor={'#ff5722'}
            circleColor={'#ff5722'}
        />
        <CircleButton
            icon={'icons/cxee.png'}
            dx={840}
            dy={1600}
            width={200}
            height={200}
            onClick={() => Music.Play("https://soundcloud.com/you/likes/song.mp3")}
        />

        <TextView style={styles.title}
                  dx={120}
                  dy={1830}
                  width={799.4}
                  height={200}
                  align={'CENTER'}
                  textSize={17}
                  textColor={'#444444'}
                  text={'Mahdi Abdolmaleki'}
        />

    </View>
);
