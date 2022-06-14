// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package com.example.ondeck.mysql

import java.sql.Connection
import java.sql.SQLException
import java.sql.Statement
import java.sql.Timestamp
import java.time.Instant

const val createCity = """-- name: createCity :exec
INSERT INTO city (
    name,
    slug
) VALUES (
    ?,
    ? 
)
"""

const val createVenue = """-- name: createVenue :execresult
INSERT INTO venue (
    slug,
    name,
    city,
    created_at,
    spotify_playlist,
    status,
    statuses,
    tags
) VALUES (
    ?,
    ?,
    ?,
    NOW(),
    ?,
    ?,
    ?,
    ?
)
"""

const val deleteVenue = """-- name: deleteVenue :exec
DELETE FROM venue
WHERE slug = ? AND slug = ?
"""

const val getCity = """-- name: getCity :one
SELECT slug, name
FROM city
WHERE slug = ?
"""

const val getVenue = """-- name: getVenue :one
SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE slug = ? AND city = ?
"""

const val listCities = """-- name: listCities :many
SELECT slug, name
FROM city
ORDER BY name
"""

const val listVenues = """-- name: listVenues :many
SELECT id, status, statuses, slug, name, city, spotify_playlist, songkick_id, tags, created_at
FROM venue
WHERE city = ?
ORDER BY name
"""

const val updateCityName = """-- name: updateCityName :exec
UPDATE city
SET name = ?
WHERE slug = ?
"""

const val updateVenueName = """-- name: updateVenueName :exec
UPDATE venue
SET name = ?
WHERE slug = ?
"""

const val venueCountByCity = """-- name: venueCountByCity :many
SELECT
    city,
    count(*)
FROM venue
GROUP BY 1
ORDER BY 1
"""

data class VenueCountByCityRow (
  val city: String,
  val count: Long
)

class QueriesImpl(private val conn: Connection) : Queries {

  @Throws(SQLException::class)
  override fun createCity(name: String, slug: String) {
    conn.prepareStatement(createCity).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun createVenue(
      slug: String,
      name: String,
      city: String,
      spotifyPlaylist: String,
      status: VenuesStatus,
      statuses: String?,
      tags: String?): Long {
    return conn.prepareStatement(createVenue, Statement.RETURN_GENERATED_KEYS).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, name)
          stmt.setString(3, city)
          stmt.setString(4, spotifyPlaylist)
          stmt.setString(5, status.value)
          stmt.setString(6, statuses)
          stmt.setString(7, tags)

      stmt.execute()

      val results = stmt.generatedKeys
      if (!results.next()) {
          throw SQLException("no generated key returned")
      }
	  results.getLong(1)
    }
  }

  @Throws(SQLException::class)
  override fun deleteVenue(slug: String, slug_2: String) {
    conn.prepareStatement(deleteVenue).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, slug_2)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun getCity(slug: String): City? {
    return conn.prepareStatement(getCity).use { stmt ->
      stmt.setString(1, slug)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = City(
                results.getString(1),
                results.getString(2)
            )
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun getVenue(slug: String, city: String): Venue? {
    return conn.prepareStatement(getVenue).use { stmt ->
      stmt.setString(1, slug)
          stmt.setString(2, city)

      val results = stmt.executeQuery()
      if (!results.next()) {
        return null
      }
      val ret = Venue(
                results.getLong(1),
                VenuesStatus.lookup(results.getString(2))!!,
                results.getString(3),
                results.getString(4),
                results.getString(5),
                results.getString(6),
                results.getString(7),
                results.getString(8),
                results.getString(9),
                results.getTimestamp(10).toInstant()
            )
      if (results.next()) {
          throw SQLException("expected one row in result set, but got many")
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun listCities(): List<City> {
    return conn.prepareStatement(listCities).use { stmt ->
      
      val results = stmt.executeQuery()
      val ret = mutableListOf<City>()
      while (results.next()) {
          ret.add(City(
                results.getString(1),
                results.getString(2)
            ))
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun listVenues(city: String): List<Venue> {
    return conn.prepareStatement(listVenues).use { stmt ->
      stmt.setString(1, city)

      val results = stmt.executeQuery()
      val ret = mutableListOf<Venue>()
      while (results.next()) {
          ret.add(Venue(
                results.getLong(1),
                VenuesStatus.lookup(results.getString(2))!!,
                results.getString(3),
                results.getString(4),
                results.getString(5),
                results.getString(6),
                results.getString(7),
                results.getString(8),
                results.getString(9),
                results.getTimestamp(10).toInstant()
            ))
      }
      ret
    }
  }

  @Throws(SQLException::class)
  override fun updateCityName(name: String, slug: String) {
    conn.prepareStatement(updateCityName).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun updateVenueName(name: String, slug: String) {
    conn.prepareStatement(updateVenueName).use { stmt ->
      stmt.setString(1, name)
          stmt.setString(2, slug)

      stmt.execute()
    }
  }

  @Throws(SQLException::class)
  override fun venueCountByCity(): List<VenueCountByCityRow> {
    return conn.prepareStatement(venueCountByCity).use { stmt ->
      
      val results = stmt.executeQuery()
      val ret = mutableListOf<VenueCountByCityRow>()
      while (results.next()) {
          ret.add(VenueCountByCityRow(
                results.getString(1),
                results.getLong(2)
            ))
      }
      ret
    }
  }

}

