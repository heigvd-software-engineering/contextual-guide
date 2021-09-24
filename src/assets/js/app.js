(function () {

    // Display the maps
    document.querySelectorAll('.map').forEach(elem => {
        new maplibregl.Map({
            container: elem,
            style: 'https://vectortiles.geo.admin.ch/styles/ch.swisstopo.leichte-basiskarte_world.vt/style.json', // stylesheet location
            center: [elem.dataset.longitude, elem.dataset.latitude], // starting position [lng, lat]
            zoom: 9 // starting zoom
        });
    });

})();
