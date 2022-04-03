import logging

from watergrid.metrics.MetricsExporter import MetricsExporter


class ConsoleExporter(MetricsExporter):
    def __init__(self):
        super().__init__()
        logging.basicConfig(level=logging.INFO)
        self._logger = logging.getLogger(__name__)

    def start_pipeline(self, pipeline_name):
        self._logger.info("Starting pipeline: %s", pipeline_name)

    def end_pipeline(self):
        self._logger.info("Ending pipeline")

    def start_step(self, step_name):
        self._logger.info("Starting step: %s", step_name)

    def end_step(self):
        self._logger.info("Ending step")

    def capture_exception(self, exception: Exception):
        self._logger.error("Exception: %s", exception)