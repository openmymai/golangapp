import React, { useState } from "react"
import useSWR from "swr"

async function fetcher(url) {
  const resp = await fetch(url)
  return resp.text()
}

function Index() {
  const { data, error } = useSWR("http://localhost:8080/api", fetcher, { refreshInterval: 1000 })

  return (
    <div>
      <h1>Hello, world!</h1>
      <p>
        This is <code>pages/index.js</code>.
      </p>
      <h2>Memory allocation stats from Go server</h2>
      {error && (
        <p>
          Error fetching profile: <strong>{error}</strong>
        </p>
      )}
      {!error && !data && <p>Loading ...</p>}
      {!error && data && <pre>{data}</pre>}

    </div>
  )
}

export default Index
