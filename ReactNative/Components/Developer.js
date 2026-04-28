import React from 'react';
import { ScrollView, Text, StyleSheet, View } from 'react-native';

export default function Developer() {
  return (
    <ScrollView style={styles.container}>
      <View style={styles.devCard}>
        <Text style={styles.devName}>Gustavo Takemura</Text>
        <Text style={styles.devSub}>Aluno 3ads</Text>
        <View style={styles.divider} />
        <Text style={styles.description}>
         Programação React Native.
        {"\n\n"}Programação Golang.
        {"\n\n"}Desenvolvimento de API REST.
        {"\n\n"}Design de UI/UX.
        </Text>
      </View>
      <View style={styles.devCard}>
        <Text style={styles.devName}>Leonardo Castelan</Text>
        <Text style={styles.devSub}>Aluno 3ads</Text>
        <View style={styles.divider} />
        <Text style={styles.description}>
          Documentação
        {"\n\n"}Criação dos slides
        {'\n\n'}Pesquisa de conteúdo
        
        </Text>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#000', padding: 20 },
  title: { color: '#00f0ff', fontSize: 22, fontWeight: 'bold', marginBottom: 20 },
  devCard: { backgroundColor: '#111', padding: 20, borderTopWidth: 5, borderTopColor: '#fcee0a' },
  devName: { color: '#fff', fontSize: 20, fontWeight: '900' },
  devSub: { color: '#ff003c', fontSize: 12, fontWeight: 'bold', marginTop: 5 },
  divider: { height: 1, backgroundColor: '#333', marginVertical: 15 },
  description: { color: '#aaa', fontSize: 14, fontStyle: 'italic' }
});