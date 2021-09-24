(function () {

    // Display the maps
    document.querySelectorAll('.map').forEach(elem => {

        // Create the map
        const map = new maplibregl.Map({
            container: elem,
            style: 'https://vectortiles.geo.admin.ch/styles/ch.swisstopo.leichte-basiskarte_world.vt/style.json',
            center: [elem.dataset.longitude, elem.dataset.latitude],
            zoom: 9 // starting zoom
        });

        // Add the marker
        new maplibregl.Marker({
            color: "#e53935",
            draggable: true
        }).setLngLat([elem.dataset.longitude, elem.dataset.latitude])
            .addTo(map);
    });

})();
