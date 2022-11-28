// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
    production: false,
    mapbox: {
        accessToken: "pk.eyJ1Ijoic2FsYXphcmJyb2NrIiwiYSI6ImNqd3hxaDBleTFobGQ0Y2sxc3lpdnNwYXMifQ.TgYFgyclmL6pGXXANhnlsw",
    },
    firebaseConfig: {
        apiKey: "AIzaSyCgE3yn1b5HkdLtBcs9U1FprH-Fs5fcCxA",
        authDomain: "cheers-a275e.firebaseapp.com",
        projectId: "cheers-a275e",
        storageBucket: "cheers-a275e.appspot.com",
        messagingSenderId: "988221557023",
        appId: "1:988221557023:web:d8111009a321d2f40a482a",
        measurementId: "G-V2X6SNDWRL",
        vapidKey: "BM86igrGAgCvUD8ZoQAkzIAdKR9nPHmVg7vIqm1weCKoWP1Y1crFwyNRgDPPRAbqetXUes4Vn4agz1gNeTF4lhM"
    },
    stripe: {
        public_key: 'pk_test_51KWqPTAga4Q2CELOu5oK8GHRPlQwVPvcISBMuoWU5yxP8VrtmBhRGm0TBKaKeKm1tz2EY7gmmvvYuFWMJEzWvFhC00qOX6gQb1'
    },
    GATEWAY_URL: "http://localhost:8080",
    // GATEWAY_URL: "http://35.230.155.254:80",
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.
