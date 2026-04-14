import React from 'react';
import { ScrollView, Text, Image, View, StyleSheet } from 'react-native';

export default function FighterDetails({ route }) {
  const { fighter } = route.params;
  return (
    <ScrollView style={styles.container}>
      <Text style={styles.title}>{fighter.nickname}</Text>
      <Image source={{ uri: fighter.image_url }} style={styles.detailImage} />
      <View style={styles.infoBox}>
        <Text style={styles.nameLabel}>{fighter.name}</Text>
        <Text style={styles.dataText}>CLASSE: {fighter.weight_class}</Text>
        <Text style={styles.dataText}>ESTILO: {fighter.style}</Text>
        <Text style={styles.dataText}>CARTEL: {fighter.wins}V - {fighter.losses}D - {fighter.draws}E</Text>
        <View style={[styles.statusTag, { backgroundColor: fighter.active ? '#00f0ff' : '#ff003c' }]}>
           <Text style={styles.statusText}>{fighter.active ? 'ATIVO' : 'APOSENTADO'}</Text>
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#000', padding: 20 },
  title: { color: '#fcee0a', fontSize: 26, fontWeight: '900', textTransform: 'uppercase', marginBottom: 10},
  detailImage: { width: '100%', height: 300, borderWidth: 2, borderColor: '#fcee0a' },
  infoBox: { marginTop: 20, padding: 15, backgroundColor: '#111', borderTopWidth: 1, borderTopColor: '#333' },
  nameLabel: { color: '#fff', fontSize: 26, fontWeight: '900', textTransform: 'uppercase', marginBottom: 10 },
  dataText: { color: '#00f0ff', fontSize: 14, marginBottom: 5, fontWeight: 'bold' },
  statusTag: { marginTop: 15, padding: 5, width: 140, alignItems: 'center' },
  statusText: { color: '#000', fontWeight: 'bold', fontSize: 10 }
});