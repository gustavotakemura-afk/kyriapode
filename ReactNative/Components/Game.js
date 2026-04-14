import React from 'react';
import { ScrollView, Text, StyleSheet, View } from 'react-native';

export default function Game() {
  return (
    <ScrollView style={styles.container}>
      <View style={styles.headerBox}>
        <Text style={styles.title}>MMA</Text>
      </View>
      <Text style={styles.bodyText}>
       O Mixed Martial Arts (Artes Marciais Mistas) é um desporto de combate de pleno contacto 
          que permite técnicas de batida (striking) e de luta no solo (grappling), combinando 
          disciplinas como Boxe, Muay Thai, Jiu-Jitsu e Wrestling.
      </Text>
      <Text style={styles.bodyText2}>
        As competições são divididas em classes de peso para garantir a equidade física entre os atletas. As principais são:
        Peso Mosca: até 56,7 kg.
        Peso Galo: até 61,2 kg.
        Peso Pena: até 65,8 kg.
        Peso Leve: até 70,3 kg.
        Peso Meio-Médio: até 77,1 kg.
        Peso Médio: até 83,9 kg.
        Peso Meio-Pesado: até 93,0 kg.
        Peso Pesado: até 120,2 kg.
      </Text>
      <Text style={styles.bodyText3}>
        Uma luta de MMA pode ser encerrada das seguintes maneiras:
        Nocaute (KO): Quando um lutador fica inconsciente devido a golpes legais.
        Nocaute Técnico (TKO): Quando o árbitro interrompe a luta por julgar que um atleta não consegue mais se defender.
        Finalização: Quando um lutador desiste via tapinha (tap out) devido a um estrangulamento ou chave de articulação.
        Decisão dos Juízes: Se a luta chegar ao fim do tempo regulamentar, três juízes pontuam o combate (geralmente usando o sistema de 10 pontos).
      </Text>
    </ScrollView>
    
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#000', padding: 20 },
  headerBox: { borderBottomWidth: 4, borderBottomColor: '#ff003c', marginBottom: 20 },
  title: { color: '#fcee0a', fontSize: 24, fontWeight: '900', fontStyle: 'italic' },
  bodyText: { color: '#fff', fontSize: 16, lineHeight: 24, textAlign: 'justify', backgroundColor: '#111', padding: 15, borderLeftWidth: 2, borderLeftColor: '#00f0ff' },
  bodyText2: { color: '#fff', fontSize: 16, lineHeight: 24, textAlign: 'justify', backgroundColor: '#111', padding: 15, borderLeftWidth: 2, borderLeftColor: '#ea00d9' },
  bodyText3: { color: '#fff', fontSize: 16, lineHeight: 24, textAlign: 'justify', backgroundColor: '#111', padding: 15, borderLeftWidth: 2, borderLeftColor: '#ff003c' },

});