import SwitchButton, {SwitchButtonValues} from "./packages/base/SwitchButton";
import Image, {ImageValues} from "./packages/base/Image";
import SliderView, {SliderViewValues} from "./packages/base/SliderView";
import TextInput from "./packages/base/TextInput";
import TextMessage from "./packages/base/TextMessage";
import ChartView from "./packages/base/ChartView";
import Button from "./packages/base/Button";
import TextBox from "./packages/base/TextBox";
import ScrollView from "./packages/base/ScrollView";

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
    switch_2.dy = 1600;
    switch_2.line = true;
    switch_2.checked = true;
    switch_2.round = 3;
    switch_2.backgroundColor = '#ff9800';

    avatar3.id = 'avatar1';
    avatar3.source = 'music/face/peoples_71.jpg';
    avatar3.width = 400;
    avatar3.height = 400;
    avatar3.dx = 340;
    avatar3.dy = 1650;
    avatar3.round = 10;

    avatar.width = 400;
    avatar.height = 400;
    avatar.round = 53;
    avatar.id = 'avatar2';
    avatar.source = 'music/face/peoples_86.jpg';
    avatar.dx = 50;
    avatar.dy = 650;


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
        <ScrollView
            dx={10}
            dy={2000}
            width={1080}
            height={600}
        />
        <TextInput
            text={"Mahdi"}
            dx={100}
            dy={200}
        />
        <TextMessage
            text={"Text Message"}
        />
        <ChartView
            id={'Mahdi 01236'}
            dx={120}
            dy={700}
            width={350}
            padding={25}
            margin={0}
            round={20}

            headerHeight={60}
            chartHeight={200}
            footerHeight={190}

            avatar={'call2/ali5.jpg'}
            icon={'icons/health_blood_pressure.png'}
            title={"Blood pressure"}

            dataArray={dataArray}
            rowArray={rowArray5}
            columnArray={columnArray}

            backgroundColor={'#ff9800'}
        />
        <SwitchButton
            id={'switchMain'}
            dx={140}
            dy={1835}
            width={170}
            height={80}
            round={12}
            checked={false}
            duration={250}
            line={true}
            backgroundColor={'#ddd'}
            activeColor={'#ff9800'}
            circleColor={'#fff'}
            onClick={changePhoto}
        />
        <Button
            dx={590}
            dy={1325}
            width={350}
            height={90}
            round={15}
            text={'Select Location'}
            align={'CENTER'}
            textSize={16}
            textColor={'#000'}
            backgroundColor={'#ddd'}
            activeColor={'#eee'}
            circleColor={'#fff'}
        />
        <TextBox
            dx={50}
            dy={4950}
            width={650}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#ddd'}
            text={'این تصویر تلسکوپی عمیق، طیف رنگ‌ها و تقارن‌های سحابی زنبق را نشان می‌دهد که در میدان‌های اطراف گرد و غبار بین ستاره‌ای جاسازی شده‌اند.'}
        />
        <TextBox
            dx={50}
            dy={5300}
            width={650}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#ddd'}
            text={'این تصویر تلسکوپی عمیق، طیف رنگ‌ها و تقارن‌های سحابی زنبق را نشان می‌دهد که در میدان‌های اطراف گرد و غبار بین ستاره‌ای جاسازی شده‌اند.این تصویر تلسکوپی عمیق، طیف رنگ‌ها و تقارن‌های سحابی زنبق را نشان می‌دهد که در میدان‌های اطراف گرد و غبار بین ستاره‌ای جاسازی شده‌اند. '}
        />
        <TextBox
            dx={250}
            dy={8550}
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
            source={'home/00/12.jpeg'}
            dx={50}
            dy={2900}
            width={800}
            height={900}
            round={10}
        />
        <Image
            source={'home/00/13.jpeg'}
            dx={50}
            dy={4000}
            width={800}
            height={800}
            round={15}
        />
        <Image
            source={'a0/012.jpg'}
            dx={50}
            dy={6000}
            width={800}
            height={800}
            round={20}
        />
        <TextBox
            dx={50}
            dy={6990}
            width={650}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={15}
            textColor={'#fff'}
            boxColor={'#0091ea'}
            text={'ما می‌توانیم در آسمان شب به تصویر مشتری و زحل اشاره کنیم و آن‌ها را «سیاره» خطاب کنیم اما تعریف این کلمه نیاز به دقت زیادی دارد. کلمه سیاره از کلمه یونانی «planēt» به معنای «سرگردان» گرفته شده است. بسیاری از منجمین باستان می‌توانستند این اجرام را که تصور می‌شد ستارگان متحرک هستند، به‌راحتی با چشم غیرمسلح ببینند.'}
        />
        <SliderView
            dx={50}
            dy={7720}
            width={990}
            height={100}
            round={4}
            line={true}
            circleSize={20}
            backgroundColor={'#f44336'}
            circleColor={'#ff9800'}
        />
        <SwitchButton
            dx={50}
            dy={8100}
            width={500}
            height={200}
            round={40}
            line={true}
            circleSize={20}
            backgroundColor={'#ddd'}
            circleColor={'#fff'}
        />
        <Image
            source={'a0/011.jpg'}
            dx={50}
            dy={9000}
            width={600}
            height={900}
            round={15}
        />
        <Image
            source={'a0/sa_13.jpg'}
            dx={50}
            dy={10300}
            width={600}
            height={900}
            round={15}
        />
        <Image
            source={'a0/sa_12.jpg'}
            dx={50}
            dy={11300}
            width={600}
            height={900}
            round={1}
        />
        <TextBox
            dx={50}
            dy={12150}
            width={520}
            round={1}
            padding={80}
            align={'NORMAL'}
            textSize={13}
            textColor={'#000'}
            boxColor={'#f5f5f5'}
            text={'                                                                  ما می‌توانیم در آسمان شب به تصویر مشتری و زحل اشاره کنیم و آن‌ها را «سیاره» خطاب کنیم اما تعریف این کلمه نیاز به دقت زیادی دارد. کلمه سیاره از کلمه یونانی «planēt» به معنای «سرگردان» گرفته شده است. بسیاری از منجمین باستان می‌توانستند این اجرام را که تصور می‌شد ستارگان متحرک هستند، به‌راحتی با چشم غیرمسلح ببینند.'}
        />

        <Button
            dx={200}
            dy={2000}
            width={600}
            height={100}

            round={50}

            //icon={'icons/health_blood_pressure.png'}

            textColor={'#fff'}
            text={'Mahdi abdolmaleki'}
            align={'CENTER'}
            textSize={20}

            backgroundColor={'#ddd'}
            activeColor={'#8d6e63'}
            circleColor={'#fff'}
        />

        <Button
            dx={100}
            dy={2300}
            width={800}
            height={120}
            round={10}

            padding={30}
            text={'+98 09125640296'}
            align={'LEFT'}
            textSize={20}

            backgroundColor={'#c8e6c9'}
            textColor={'#2e7d32'}
            activeColor={'#e8f5e9'}
            circleColor={'#fff'}
        />
        <Button
            dx={200}
            dy={2600}
            width={200}
            height={200}
            round={100}

            icon={'icons/health_blood_pressure.png'}
            textColor={'#fff'}
            text={''}
            align={'CENTER'}
            textSize={20}

            backgroundColor={'#ddd'}
            activeColor={'#ffcdd2'}
            circleColor={'#fff'}
        />

    </View>
);
