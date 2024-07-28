import Image from "./packages/base/Image";
import TextView from "./packages/base/TextView";

const ViewPager = () => (

    <View>
        <Image
            source={'posters/thumbnail_2/2lAPv2jQ5eh54l4EttA9wjDljED.jpg'}
            dx={50}
            dy={200}
            width={300}
            height={500}
        />
        <Image
            source={'posters/thumbnail_2/2XcN2aIma8TqShsOBPbcpxLzY1A.jpg'}
            dx={400}
            dy={200}
            width={300}
            height={500}
        />
        <Image
            source={'posters/thumbnail_2/y9DdvY1IRNS2hWLq161hmP2gXn.jpg'}
            dx={750}
            dy={200}
            width={300}
            height={500}
        />
        <TextView
            dx={50}
            dy={720}
            width={300}
            textColor={'#000'}
            textSize={15}
            text={"Dear Jon"}
            align={"CENTER"}

        />
        <TextView
            dx={400}
            dy={720}
            width={300}
            text={"Family Troop"}
            textColor={'#000'}
            textSize={15}
            align={"CENTER"}
        />
        <TextView
            dx={750}
            dy={720}
            width={300}
            text={"Beautiful Live"}
            textColor={'#000'}
            textSize={15}
            align={"CENTER"}
        />
    </View>
);
