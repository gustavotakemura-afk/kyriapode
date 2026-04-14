import React, { useEffect, useState } from 'react';
import { View, Text, FlatList, Image, TouchableOpacity, StyleSheet } from 'react-native';

export default function FightersList({ navigation }) {
  const [fighters, setFighters] = useState([]);

  useEffect(() => {
      fetch('http://localhost:8080/fighters')
      .then(res => res.json())
      .then(data => setFighters(data))
      .catch(err => console.error('Erro ao buscar lutadores:', err));
  }, []);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>LUTADORES</Text>
      
      <FlatList
        data={fighters}
        keyExtractor={item => item.id.toString()}
        renderItem={({ item }) => (
          <TouchableOpacity
            style={styles.card}
            onPress={() => navigation.navigate('FighterDetails', { fighter: item })}
          >
            <View style={{ marginLeft: 15 }}>
              <Text style={styles.name}>{item.name}</Text>
              <Text style={styles.nickname}>"{item.nickname}"</Text>
              
              <View style={[
                styles.glitchLine, 
                { backgroundColor: item.active ? '#00ff41' : '#ff003c' }
              ]} />
              
              <Text style={styles.statusLabel}>
                {item.active ? 'ATIVO' : 'APOSENTADO'}
              </Text>
            </View>
          </TouchableOpacity>
        )}
      />

      <TouchableOpacity
        style={styles.buttonMain}
        onPress={() => navigation.navigate('Game')}
      >
        <Text style={styles.buttonText}>Regras</Text>
      </TouchableOpacity>

      <TouchableOpacity
        style={styles.buttonSecondary}
        onPress={() => navigation.navigate('Developer')}
      >
        <Text style={styles.buttonSecondaryText}>CREDITOS</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { 
    flex: 1, 
    backgroundColor: '#000', 
    padding: 20, 
    paddingTop: 60
  },
  title: { 
    fontSize: 32, 
    fontWeight: '900', 
    color: '#fcee0a', 
    fontStyle: 'italic', 
    marginBottom: 30, 
    textTransform: 'uppercase' 
  },
  card: { 
    flexDirection: 'row', 
    backgroundColor: '#0d0d0d', 
    marginBottom: 15, 
    padding: 15, 
    borderLeftWidth: 4, 
    borderLeftColor: '#fcee0a', 
    alignItems: 'center' 
  },
  image: {  
    height: 20,
    width: 20
  },
  name: { 
    color: '#fff', 
    fontSize: 20, 
    fontWeight: 'bold', 
    textTransform: 'uppercase' 
  },
  nickname: { 
    color: '#aaa', 
    fontSize: 14 
  },
  glitchLine: { 
    height: 3, 
    width: 80, 
    marginTop: 8 
  },
  statusLabel: { 
    fontSize: 9, 
    color: '#444', 
    marginTop: 4, 
    fontWeight: '900' 
  },

  buttonMain: { 
    marginTop: 10,
    padding: 10,
    borderWidth: 1,
    borderColor: '#fcee0a',
    alignItems: 'flex-start',
    backgroundColor: 'transparent'
  },
  buttonText: { 
    color: '#fcee0a',
    fontSize: 10,
    fontWeight: 'bold',
    textTransform: 'uppercase'
  },

  buttonSecondary: {
    marginTop: 10,
    padding: 10,
    borderWidth: 1,
    borderColor: '#ff003c',
    alignItems: 'flex-start',
    backgroundColor: 'transparent'
  },
  buttonSecondaryText: {
    color: '#ff003c',
    fontSize: 10,
    fontWeight: 'bold',
    textTransform: 'uppercase'
  }
});