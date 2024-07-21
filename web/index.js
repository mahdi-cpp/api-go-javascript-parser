import SwitchButton, {SwitchButtonValues} from "./component/SwitchButton";
import Image, {ImageValues} from "./component/Image";

function changePhoto() {

    let switch1 = SwitchButtonValues
    let switch2 = SwitchButtonValues

    switch1.id = 'switch1'
    switch1.dx = 45;

    switch2.id = 'switch2'
    switch2.dx = 500;

    let imageTest = ImageValues
    imageTest.id = 'image1'
    imageTest.source = 'call2/ali4.jpg'

    switch1.checked = true;
    switch1.line = true
}

function changePhoto2() {
    let image1 = ImageValues
    image1.id = 'image1'
    image1.source = 'call2/ali4.jpg'
}

const Index = () => (
    <View>
        <SwitchButton
            onClick={startAnimation}
            onPinch={}
        />
        <SwitchButton
            id={'switch1'}
            dx={50}
            dy={460}
            width={280}
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
            dx={850}
            dy={260}
            width={180}
            height={90}
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
            id={"image1"}
            source={'call2/ali5.jpg'}
            dx={50}
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
            width={960}
            height={90}
            round={15}
            circleSize={10}
            line={true}
            backgroundColor={'#fbe9e7'}
            activeColor={'#ff5722'}
            circleColor={'#ff5722'}
        />
        <CircleButton
            icon={'icons/cxee.png'}
            dx={440}
            dy={1600}
            width={200}
            height={200}
            onClick={() => Music.Play("https://soundcloud.com/you/likes/song.mp3")}
        />
        <Image
            style={styles.image}
            width={500}
            height={950}
            dx={300}
            dy={610}
            round={15}
            source={'messi/2017-07-29_00-3-42_UTC_1.jpg'}
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
