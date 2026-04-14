// App.js
import React, { useEffect, useState } from 'react';
import { View, Text, FlatList, Image, TouchableOpacity, StyleSheet, ScrollView } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';


import FightersList from './Components/FightersList';
import FighterDetails from './Components/FightersDetail';
import Game from './Components/Game';
import Developer from './Components/Developer';

const Stack = createNativeStackNavigator();





export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator screenOptions={{ 
        headerStyle: { backgroundColor: '#000' },
        headerTintColor: '#fcee0a',               
        headerTitleStyle: { fontWeight: '900', textTransform: 'uppercase' },
        headerShadowVisible: false,               
      }}>
        <Stack.Screen name="FightersList" component={FightersList} options={{ title: 'Lutadores' }} />
        <Stack.Screen name="FighterDetails" component={FighterDetails} options={{ title: 'Detalhes' }} />
        <Stack.Screen name="Game" component={Game} options={{ title: 'Regras' }} />
        <Stack.Screen name="Developer" component={Developer} options={{ title: 'Sobre o Dev' }} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
