import { useState } from "react";

const API = import.meta.env.VITE_API_URL as string;

export default function App() {
  const [input, setInput] = useState("");
  const [msg, setMsg] = useState("");

  const submit = async () => {
    if (!input.trim()) {
      setMsg("input required");
      return;
    }

    setMsg("sending...");

    try {
      const res = await fetch(`${API}/users`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ input })
      });

      const data = await res.json();

      if (!res.ok) throw new Error(data.error || "error");

      setMsg("success");
      setInput("");
    } catch (e: any) {
      setMsg(e.message);
    }
  };

  return (
    <div style={wrap}>
      <h2>User Parser</h2>

      <input
        style={inputStyle}
        placeholder="CUT MINI 28 BANDA ACEH"
        value={input}
        onChange={(e) => setInput(e.target.value)}
      />

      <button style={btn} onClick={submit}>
        Submit
      </button>

      {msg && <p>{msg}</p>}
    </div>
  );
}

const wrap: React.CSSProperties = {
  maxWidth: 800,
  display: "flex",
  flexDirection: "column",
  position: "fixed",
  left: "40%",
  top: "20%",
  gap: 12
};

const inputStyle: React.CSSProperties = {
  padding: 10
};

const btn: React.CSSProperties = {
  padding: 10,
  cursor: "pointer"
};