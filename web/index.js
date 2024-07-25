import SwitchButton, {SwitchButtonValues} from "./SwitchButton";
import Image, {ImageValues} from "./Image";

const SwitchButtonValues: SwitchButtonValues = {};
const ImageValues: ImageValues = {}


function changePhoto() {
    let switch_1 = SwitchButtonValues
    let switch_2 = SwitchButtonValues
    let avatar = ImageValues
    let avatar3 = ImageValues


    switch_1.id = 'switch1';
    switch_1.dx = 545;
    switch_1.dy = 600;
    switch_1.width = 400
    switch_1.height = 150;
    switch_1.round = 15;
    switch_1.icon = 'icons/download_80_1.png'
    switch_1.checked = true;
    switch_1.line = true;
    switch_1.activeColor = '#ff9800'
    switch_1.backgroundColor = '#888';

    switch_2.id = 'switch2'; // Declare x, give it the value of 5
    switch_2.icon = 'icons/cast.png'
    switch_2.width = 400;
    switch_2.height = 200
    switch_2.dx = 600// Declare x, give it the value of 5
    switch_2.dy = 1200;
    switch_2.line = true;
    switch_2.checked = true;
    switch_2.round = 3;
    switch_2.backgroundColor = '#ff9800';

    avatar3.id = 'avatar1';
    avatar3.source = 'music/face/peoples_73.jpg';
    avatar3.width = 400;
    avatar3.height = 400;

    avatar.width = 400;
    avatar.height = 400;
    avatar.round = 53;
    avatar.id = 'avatar2';
    avatar.source = 'music/face/peoples_86.jpg';
    avatar.dx = 50;
    avatar.dy = 250;

    avatar3.dx = 150;
    avatar3.dy = 1350;
    avatar3.round = 10;
}

function changePhoto2() {
    let image1 = ImageValues
    let image2 = ImageValues

    image1.id = 'image1'
    image2.id = 'image2'

    image1.dx = 200
    image2.dy = 300
}

function musicPlay(url, count) {
    /*
dsfd.id = 'switch2';
gfgh if comment
ali.id = 'swihghgtch2';
*/
    let music1 = ImageValues
    /*
    dsfd.id = 'switch2';
    gfgh if comment
    ali.id = 'swihghgtch2';
    */
    music1.id = 'musicPlayer'
    music1.source = "music/track/artist1.mp4"
    /*
    dsfd.id = 'switch2';
    gfgh if comment
    ali.id = 'swihghgtch2';
    */
}

const Index = () => (
    <View>
        <SwitchButton
            onClick={changePhoto2}
            onPinch={}
        />
        <SwitchButton
            id={'switchMain'}
            dx={630}
            dy={300}
            width={400}
            height={200}
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
            id={'switch1'}
            dx={720}
            dy={600}
            width={200}
            height={90}
            round={5}
            checked={false}
            duration={250}
            line={true}
            backgroundColor={'#eceff1'}
            activeColor={'#ff9800'}
            circleColor={'#fff'}
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
            id={"avatar1"}
            source={'call2/ali5.jpg'}
            dx={50}
            dy={310}
            width={170}
            height={170}
            round={55}
            onClick={changePhoto}
        />
        <Image
            id={"avatar2"}
            source={'call2/ali4.jpg'}
            dx={250}
            dy={310}
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
