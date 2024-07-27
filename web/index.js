import SwitchButton, {SwitchButtonValues} from "./component/SwitchButton";
import Image, {ImageValues} from "./component/Image";
import {SliderViewValues} from "./component/SliderView";
import ChartView from "./component/ChartView";
import TextBox from "./component/TextBox";

const SwitchButtonValues: SwitchButtonValues = {};
const ImageValues: ImageValues = {}
const SliderViewValues: SliderViewValues = {}

// // Create an object that implements the Person interface
// let person: ChartViewProps = {
//     caption: "Alice",
//     verticallyArray: [25, 30, 35]
// };

function changePhoto() {

    let switch_1 = SwitchButtonValues
    let switch_2 = SwitchButtonValues
    let avatar = ImageValues
    let avatar3 = ImageValues
    let volume = SliderViewValues

    volume.id = 'Volume'
    volume.height = 50
    volume.round = 5
    volume.dy = 950
    volume.activeColor = '#ff9800'
    volume.circleColor = '#ff9800'

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

function messageSend() {

    let avatarAbedNaseri = ImageValues
    let avatarFarahmand = ImageValues

    avatarFarahmand.dx = 150;
    avatarFarahmand.dy = 500;
    avatarFarahmand.width = 300;
    avatarFarahmand.height = 300;
    avatarFarahmand.round = 10;
    avatarFarahmand.id = 'avatar1';
    avatarFarahmand.source = 'music/nf1/bb25.jpg';

    avatarAbedNaseri.dx = 550;
    avatarAbedNaseri.dy = 500;
    avatarAbedNaseri.width = 300;
    avatarAbedNaseri.height = 300;
    avatarAbedNaseri.round = 10;
    avatarAbedNaseri.id = 'avatar2';
    avatarAbedNaseri.source = 'music/nf1/bb21.jpg';
}

const Index = () => (
    <View>
        <ChartView
            dx={10}
            dy={900}
            width={450}
            padding={25}
            margin={0}
            round={0}

            avatar={'call2/ali5.jpg'}
            icon={'icons/health_blood_pressure.png'}
            title={"Blood pressure"}

            headerHeight={60}
            chartHeight={250}
            footerHeight={180}

            rowArray={rowArray}
            columnArray={columnArray}
        />

        <TextBox
            dx={250}
            dy={2550}
            width={600}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={14}
            textColor={'#fff'}
            boxColor={'#88f'}
            text={'این تصویر تلسکوپی عمیق، طیف رنگ‌ها و تقارن‌های سحابی زنبق را نشان می‌دهد که در میدان‌های اطراف گرد و غبار بین ستاره‌ای جاسازی شده‌اند.'}
        />
        <Image
            source={'sabihe/sa_2.jpg'}
            dx={180}
            dy={110}
            width={500}
            height={700}
            round={10}
        />
        <Image
            id={"avatar1"}
            source={'call2/ali1.jpg'}
            dx={50}
            dy={715}
            width={100}
            height={100}
            round={55}
        />
        <CircleButton
            dx={0}
            dy={900}
            width={1080}
            height={1200}
        />
        <SwitchButton
            id={'switchMain'}
            dx={830}
            dy={300}
            width={200}
            height={100}
            round={5}
            checked={false}
            duration={250}
            line={true}
            backgroundColor={'#eceff1'}
            activeColor={'#00e676'}
            circleColor={'#fff'}
            onClick={changePhoto}
        />
        <CircleButton
            dx={0}
            dy={900}
            width={1080}
            height={1200}
        />

    </View>


);
