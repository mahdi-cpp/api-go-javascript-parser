function FloatView(props) {
    return null;
}

function CircleButton(props) {
    return null;
}

let screenWith = 1080;
let screenHeight;
let Music;

Music.Start = function () {

};

Music.Next = function () {

};

Music.Forward = function () {

};

Music.Play = function (mp3) {

};


const FloatView = () => (

    <FloatView>
        <View>
            <Text style={styles.title}
                  width={100}
                  align={'NORMAL'}
                  text={'Podcasts '}
            />
            <Image
                style={styles.image}
                width={200}
                height={200}
                dx={150}
                dy={50}
                source={'call2/ali4.jpg'}
            />
            <CircleButton
                icon={'icons/forward'}
                width={screenWith / 2}
                height={screenHeight}
                dx={60}
                dy={70}
                onPress={Music.Next()}
                onClick={() => Music.Play("https://soundcloud.com/you/likes/song.mp3")}
            />
            <Image
                style={styles.image}
                width={200}
                height={200}
                dx={150}
                dy={50}
                source={'call2/ali5.jpg'}
            />
            <Text style={styles.title}
                  width={100}
                  align={'NORMAL'}
                  text={'The title and new helium handler are required. It is recommended to'}
            />
            <CircleButton
                icon={'icons/play'}
                width={screenWith / 4}
                height={screenHeight * 2}
                dx={50.6}
                dy={45.5}
                onPress={Music.Play()}
                onClick={() => Music.Play("https://soundcloud.com/you/likes/song.mp3")}
            />
            <Image
                style={styles.tinyLogo}
                source={'@expo/snack-static/react-native-logo.png'}
            />
            <CircleButton
                icon={'icons/forward'}
                width={screenWith / 2}
                height={14}
                dx={20}
                dy={14}
                onPress={() => Music.Start()}
            />
            <CircleButton
                icon={'icons/forward'}
                width={screenWith / 2}
                height={14}
                dx={20}
                dy={14}
                onPress={() => Music.Next()}
            />
            <Image
                style={styles.tinyLogo}
                source={'@expo/snack-static/react-native-logo.png'}
            />
            <Text style={styles.title}
                  text={'Ali Artist new the Perimeter method for Rectangle'}
                  width={100}
                  align={'NORMAL'}
            />
        </View>
    </FloatView>
);
