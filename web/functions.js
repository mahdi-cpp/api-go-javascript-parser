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

function SwitchButton(props) {
    return null;
}
function VideoPlay(url) {
    Video.Play(url)
};

function MusicPlay(url) {
    Music.Play(url)
};
