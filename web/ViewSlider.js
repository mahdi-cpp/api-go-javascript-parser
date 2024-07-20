function MusicPlay(url) {
    Music.Play(url)
};

const ViewSlider = () => (
    <View>
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
                  width={700}
                  height={300}
                  align={'NORMAL'}
                  textSize={14}
                  textColor={'#888'}
                  text={'09125640293'}
        />

        <SwitchButton
            dx={300}
            dy={800}
            width={180}
            height={90}
            round={5}
            checked={false}
            duration={250}
            line={false}
            backgroundColor={'#eceff1'}
            activeColor={'#f44336'}
            circleColor={'#fff'}

        />

        <SwitchButton
            dx={700}
            dy={300}
            width={170}
            height={90}
            round={100}
            checked={false}
            duration={200}
            line={false}
            backgroundColor={'#cfd8dc'}
            activeColor={'#4caf50'}
            circleColor={'#fff'}
        />
        <SwitchButton
            icon={'icons/nav_check.png'}
            dx={130}
            dy={950}
            width={400}
            height={150}
            round={5}
            checked={false}
            duration={300}
            line={true}
            backgroundColor={'#b2ebf2'}
            activeColor={'#00bcd4'}
            circleColor={'#fff'}
        />
            <SwitchButton
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
                  dx={75.8}
                  dy={450}
                  width={799.4}
                  height={200}
                  align={'CENTER'}
                  textSize={17}
                  textColor={'#444444'}
                  text={'Mahdi Abdolmaleki'}
        />
        <CircleButton
            icon={'icons/tj.png'}
            dx={50}
            dy={1800}
            width={100}
            height={100}
        />
    </View>
);
