+--------------------+                           +-------------------+
|                    |                           |                   |
|   User's Browser   |                           |  Your Application |
|                    |                           |     Backend       |
+---------+----------+                           +---------+---------+
          |                                              |
          |                                              |
          | 1. User clicks "Login with Google"           |
          +--------------------------------------------> |
          |                                              |
          |                                              |
          |                                              |
          | 2. Redirect to Google's OAuth 2.0           |
          |    Authorization Endpoint                   |
          +--------------------------------------------> |
          |                                              |
          +----------------------------------+           |
          |                                  |           |
          | 3. User Grants Permission        |           |
          |    & Auth Code is returned       |           |
          |<-------------------------------->|           |
          |                                  |           |
          |                                  +----+      |
          |                                       |      |
          |                                       |      |
          | 4. Authorization Code                 |      |
          |    is sent to the Backend             |      |
          +-------------------------------------->|      |
          |                                              |
          |                                              |
          | 5. Backend sends Authorization               |
          |    Code to Google's Token Endpoint           |
          |    to exchange for Access Token              |
          +--------------------------------------------> |
          |                                              |
          |                                              |
          |                                              |
          | 6. Google responds with Access Token         |
          |    (and optionally a Refresh Token)          |
          |<-------------------------------------------- |
          |                                              |
          |                                              |
          |                                              |
          | 7. Access Token is used to access            |
          |    protected resources or APIs on behalf     |
          |    of the user                               |
          +--------------------------------------------> |
          |                                              |
          |                                              |
          |                                              |
+---------+----------+                           +---------+---------+
|                    |                           |                   |
|   Google APIs      |                           |  Your Application |
|                    |                           |     Backend       |
+--------------------+                           +-------------------+

