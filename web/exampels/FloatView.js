import {SwitchButton} from "../component/SwitchButton";

function FloatView(props) {
    return null;
}

function CircleButton(props) {
    return null;
}

let screenWith = 1080;
let screenHeight;
let Music;
let Video;

Music.Start = function () {

};

Music.Next = function () {

};

Music.Forward = function () {

};

function serverRequest(url) {
    http.get(url)
        .then((response) => {
            const data = response.data;

        })
        .catch((response) => {});
}

function TextView(props) {
    return null;
}

function VideoPlay(url) {
    Video.Play(url)
};

function MusicPlay(url) {
    Music.Play(url)
};



const FloatView = () => (

    <FloatView>
        <View>
            <SwitchButton
                dx={900}
                dy={200}
                width={120}
                height={40}
                value={0}
                android={}

            />
            <CircleButton
                icon={'icons/cxee.png'}
                dx={440}
                dy={1600}
                width={200}
                height={200}
                onClick={() => MusicPlay("https://soundcloud.com/you/likes/song.mp3")}
            />
            <Image
                style={styles.image}
                width={170}
                height={170}
                dx={50}
                dy={150}
                round={55}
                source={'call2/ali5.jpg'}
            />
            <TextView style={styles.title}
                      dx={240}
                      dy={180}
                      width={500}
                      height={200}
                      textSize={18}
                      textColor={'#000000'}
                      align={'NORMAL'}
                      text={'Farahmand Alipor'}
            />
            <TextView style={styles.title}
                      dx={240}
                      dy={250}
                      width={500}
                      height={200}
                      align={'NORMAL'}
                      textSize={14}
                      textColor={'#888'}
                      text={'09125640293'}
            />


            <Image
                style={styles.image}
                width={500}
                height={950}
                dx={300}
                dy={610}
                round={15}
                source={'messi/2017-07-29_00-30-42_UTC_1.jpg'}
            />
            <CircleButton
                icon={'icons/tj.png'}
                dx={500}
                dy={480}
                width={100}
                height={100}
            />
            <CircleButton
                icon={'icons/cxee.png'}
                dx={440}
                dy={1600}
                width={200}
                height={200}
                onPress={Music.Play()}
                onClick={() => Music.Play("https://soundcloud.com/you/likes/song.mp3")}
            />
            <CircleButton
                icon={'icons/music.png'}
                dx={870}
                dy={1900}
                width={150}
                height={150}
                onClick={() => MusicPlay("http://server/api/music/v1/artists")}
            />


            <TextView style={styles.title}
                      dx={0}
                      dy={1820}
                      width={1080}
                      height={200}
                      align={'CENTER'}
                      textSize={17}
                      textColor={'#444444'}
                      text={'Mahdi Abdolmaleki'}
            />
        </View>
    </FloatView>
);
