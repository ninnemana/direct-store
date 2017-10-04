module.exports = {
	webpack: function(config, { dev }) {
		if (config.resolve.alias) {
			delete config.resolve.alias.react
			delete config.resolve.alias['react-dom']
		}

		return config;
	}
};
