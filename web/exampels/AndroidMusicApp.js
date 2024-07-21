import React from 'react';
import {Alert, Button, StyleSheet, View,} from 'react-native';

const Separator = () => <View style={styles.separator} />;


const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        marginHorizontal: 16,
    },
    image: {
        width: '200',
        height: '200',
        dx: '150',
        dy: '50',
    },
    title: {
        textAlign: 'center',
        marginVertical: 8,
    },
    fixToText: {
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    separator: {
        marginVertical: 8,
        borderBottomColor: '#737373',
        borderBottomWidth: StyleSheet.hairlineWidth,
    },
});


const App = () => (
    <View>

        <MenuView
            width={200}
            height={400}
        />
        <FloatView
            width={screenWith}
            height={screenHeight}
        />

        <Button
            title="Press me"
            width={12}
            height={14}
            dx={20}
            dy={14}
            onPress={() => Alert.alert('Simple Button pressed')}
        />
        <Image
            style={styles.tinyLogo}
            source={'@expo/snack-static/react-native-logo.png'}
        />

    </View>
);

export default App;