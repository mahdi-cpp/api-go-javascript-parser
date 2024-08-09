import SwitchButton, {SwitchButtonValues} from "./packages/base/SwitchButton";
import Image, {ImageValues} from "./packages/base/Image";
import SliderView, {SliderViewValues} from "./packages/base/SliderView";
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
            width={900}
            height={350}
        />
        <Button
            dx={0}
            dy={0}
            width={900}
            height={300}
            round={100}
            padding={30}
            text={'Phone:  +98 09125640296'}
            align={'CENTER'}
            textSize={20}
            textColor={'#000'}

            backgroundColor={'#eee'}
            activeColor={'#eee'}
        />
        <Button
            dx={40}
            dy={20}
            width={300}
            height={300}
            padding={20}
            round={50}
            icon={'icons/health_blood_pressure.png'}
            textColor={'#000'}
            text={''}
            align={'CENTER'}
            textSize={18}
            backgroundColor={'#ddd'}
            activeColor={'#fff'}
            circleColor={'#fff'}
        />
        <Button
            dx={30}
            dy={0}
            width={800}
            height={300}
            round={7}
            padding={30}
            text={'Following'}
            align={'CENTER'}
            textSize={15}
            textColor={'#000'}
            activeColor={'#eee'}
        />
        <Button
            dx={50}
            dy={0}
            width={800}
            height={300}
            round={10}
            padding={30}
            text={'Mahdi abdolmaleki'}
            align={'CENTER'}
            textSize={15}
            textColor={'#000'}
            activeColor={'#ddd'}
        />
        <SwitchButton
            dx={50}
            dy={0}
            width={200}
            height={100}
            round={2}
            duration={300}
            line={true}
            circleSize={20}
            backgroundColor={'#eee'}
            activeColor={'#ff9800'}
            circleColor={'#fff'}
        />
        <SwitchButton
            dx={200}
            dy={0}
            width={250}
            height={100}
            round={40}
            duration={300}
            line={true}
            circleSize={20}
            backgroundColor={'#ddd'}
            activeColor={'#00e676'}
            circleColor={'#fff'}
        />
        <SwitchButton
            dx={200}
            dy={0}
            width={700}
            height={200}
            round={50}
            duration={300}
            line={true}
            circleSize={52}
            backgroundColor={'#eee'}
            activeColor={'#4fc3f7'}
            circleColor={'#fff'}
        />

        <ChartView
            id={'Mahdi 01236'}
            dx={40}
            width={900}
            padding={15}
            margin={10}
            round={20}
            headerHeight={60}
            chartHeight={200}
            footerHeight={350}

            avatar={'call2/ali5.jpg'}
            icon={'icons/health_blood_pressure.png'}
            title={"Blood pressure"}

            dataArray={dataArray}
            rowArray={rowArray5}
            columnArray={columnArray}

            backgroundColor={'#ff9800'}
        />
        <TextBox
            dx={20}
            width={330}
            round={0}
            padding={20}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'salam mahdi! rrr'}
        />
        <TextBox
            dx={20}
            width={800}
            round={0}
            padding={20}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'سلام خوبی؟ حالت چطوره؟؟!!! بهتر شدی.\nشماره فرشاد : 0912563236\n@helium'}
        />
        <TextBox
            width={800}
            padding={20}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'Oppenheimer is a 2023 epic biographical thriller drama film[a] written, directed, and produced by Christopher Nolan.[8] It follows the life of J. Robert Oppenheimer, the American theoretical physicist who helped develop the first nuclear weapons during World War II. Based on the 2005 biography American Prometheus by Kai Bird and Martin J. Sherwin, the film chronicles Oppenheimer\'s studies, his direction of the Los Alamos Laboratory and his 1954 security hearing. Cillian Murphy stars as Oppenheimer, alongside Robert Downey Jr. as the United States Atomic Energy Commission member Lewis Strauss. The ensemble supporting cast includes Emily Blunt, Matt Damon, Florence Pugh, Josh Hartnett, Casey Affleck, Rami Malek, and Kenneth Branagh.\n' +
                '\n' +
                'Oppenheimer was announced in September 2021. It is Nolan\'s first film not distributed by Warner Bros. Pictures since Memento (2000), due to his conflicts regarding the studio\'s simultaneous theatrical and HBO Max release schedule.[9] Murphy was the first cast member to sign on the following month, with the rest joining between November 2021 and April 2022. Pre-production began by January 2022, and filming took place from February to May. The cinematographer, Hoyte van Hoytema, used a combination of IMAX 65 mm and 65 mm large-format film, including, for the first time, scenes in IMAX black-and-white film photography. As with many of his previous films, Nolan used extensive practical effects, with minimal compositing.\n' +
                '\n' +
                'Oppenheimer premiered at Le Grand Rex in Paris on July 11, 2023, and was theatrically released in the United States and the United Kingdom ten days later by Universal Pictures. Its concurrent release with Warner Bros.\'s Barbie was the catalyst of the "Barbenheimer" phenomenon, encouraging audiences to see both films as a double feature. Oppenheimer grossed over US$977 million worldwide, becoming the third-highest-grossing film of 2023, the highest-grossing World War II-related film, the highest-grossing biographical film and the second-highest-grossing R-rated film.'}
        />
        <TextBox
            dx={40}
            width={800}
            round={4}
            padding={10}
            align={'NORMAL'}
            textSize={16}
            textColor={'#fff'}
            boxColor={'#fff'}
            text={'benyamin_macro.1987 Hello friends Digikala Links:\nhttps://www.digikala.com/main/book-and-media/ \n\nPlease set name and phone number \uD83D\uDE01 \uD83C\uDF52 \uD83C\uDF54 \n https://www.digikala.com/product/dkp-6396178/ \n\n 09125640236 \n\n #Englis_Learning'}
        />
        <TextBox
            dx={40}
            dy={0}
            width={800}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={16}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'این تصویر تلسکوپی عمیق، طیف رنگ‌ها و تقارن‌های سحابی زنبق را نشان می‌دهد که در میدان‌های اطراف گرد و غبار بین ستاره‌ای جاسازی شده‌اند.'}
        />
        <TextBox
            dx={40}
            dy={0}
            width={800}
            round={4}
            padding={80}
            align={'NORMAL'}
            textSize={18}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'Vegas in the palm of your hand! @ok_design_sf showed us their 75mm MiniSphere during the chaotic teardown at Open Sauce. With 1080 LEDs and an IMU, the compact design features 10 flexible circuit boards, mounted on a 3D-printed geodesic sphere housing. The nRF52-based Qorvo module provides BLE, plus an ultra-wideband radio so that multiple ‘Spheres (about twelve are planned) can coordinate patterns with each other. Vegas in the palm of your hand! @ok_design_sf showed us their 75mm MiniSphere during the chaotic teardown at Open Sauce. With 1080 LEDs and an IMU, the compact design features 10 flexible circuit boards, mounted on a 3D-printed geodesic sphere housing. The nRF52-based Qorvo module provides BLE, plus an ultra-wideband radio so that multiple ‘Spheres (about twelve are planned) can coordinate patterns with each other.'}
        />
        <TextBox
            dx={40}
            dy={0}
            width={800}
            round={1}
            padding={80}
            align={'NORMAL'}
            textSize={18}
            textColor={'#fff'}
            boxColor={'#fff'}
            text={'سلام مهدی. خوبی؟ چه خبر؟'}
        />
        <TextBox
            dx={40}
            width={800}
            round={1}
            padding={80}
            align={'NORMAL'}
            textSize={13}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'                                                           این اجرام را که تصور می‌شد ستارگان متحرک هستند، به‌راحتی با چشم غیرمسلح ب بینند.'}
        />
        <TextBox
            dx={40}
            width={800}
            round={1}
            padding={80}
            align={'NORMAL'}
            textSize={18}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'Use Messages on your iPhone or iPad Send texts, photos, videos, and more. Pin your conversations. Set Messages to automatically delete verification codes once you use them with AutoFill.'}
        />
        <TextBox
            dx={20}
            width={330}
            round={0}
            padding={20}
            align={'NORMAL'}
            textSize={15}
            textColor={'#000'}
            boxColor={'#fff'}
            text={'salam mahdi! rrr'}
        />
        <Image
            source={'tiny/13.jpg'}
            dx={50}
            width={900}
            height={800}
        />
        <Image
            source={'tiny/5.jpg'}
            dx={50}
            width={900}
            height={600}
        />
        <Image
            source={'tiny/12.jpg'}
            dx={50}
            width={900}
            height={800}
        />
        <Image
            source={'tiny/10.jpg'}
            dx={50}
            width={900}
            height={800}
        />

        <SliderView
            dx={50}
            dy={0}
            width={880}
            height={150}
            round={40}
            line={true}
            circleSize={-30}
            backgroundColor={'#fff'}
            activeColor={'#ff9800'}
            circleColor={'#ff9800'}
        />
        <SliderView
            dx={50}
            dy={0}
            width={880}
            height={10}
            round={100}
            line={true}
            circleSize={18}
            backgroundColor={'#fff'}
            activeColor={'#d4e157'}
            circleColor={'#827717'}
        />
        <SliderView
            dx={50}
            dy={0}
            width={850}
            height={100}
            round={4}
            line={true}
            circleSize={20}
            backgroundColor={'#fff'}
            activeColor={'#ff0000'}
            circleColor={'#ff0000'}
        />
        <Image
            source={''}
            dx={50}
            dy={0}
            width={900}
            height={900}
            round={15}
        />

    </View>
);
