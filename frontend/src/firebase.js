import { initializeApp } from "firebase/app";
import { getFirestore } from "firebase/firestore";

// Your Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyBvDM4XNtroMROae0y5Mnmh3PfGyt8_Ah8",
  authDomain: "mlb-predictor-448606.firebaseapp.com",
  projectId: "mlb-predictor-448606",
  storageBucket: "mlb-predictor-448606.firebasestorage.app",
  messagingSenderId: "798993946741",
  appId: "1:798993946741:web:d3ab44c10bfc5c3966403c",
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const db = getFirestore(app);
